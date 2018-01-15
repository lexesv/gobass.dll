// player
package bass

import (
	"errors"
	"log"
	"os"
	"regexp"
)

type Player struct {
	Conf    PlayerConf
	Channel int
}

type PlayerConf struct {
	Device int
	Freq   int
	Flags  int
	Vol    float32 //default 55%
	ChVol  float32 //default 75%
	Source string
}

// NewPlayer
func NewPlayer(conf PlayerConf) (*Player, error) {
	if conf.Device == 0 {
		conf.Device = -1
	}
	if conf.Vol == 0.0 {
		conf.Vol = 55
	}
	if conf.ChVol == 0.0 {
		conf.ChVol = 75
	}
	if conf.Freq == 0 {
		conf.Freq = 44100
	}
	if conf.Source == "" {
		return nil, errors.New("Source is empty")
	}
	if ok, err := Init(conf.Device, conf.Freq, conf.Flags); ok {
		return &Player{Conf: conf}, nil
	} else {
		return nil, err
	}

}

func (p *Player) Free() (err error) {
	if _, err = Free(); err != nil {
		return err
	}
	return nil
}

// Player.Play
func (p *Player) Play() (err error) {
	var ch int
	r := regexp.MustCompile(`(?isU)^http[s]?://`)
	if r.MatchString(p.Conf.Source) {
		if ch, err = StreamCreateURL(p.Conf.Source); err != nil {
			return err
		}
	} else {
		if _, err = os.Stat(p.Conf.Source); err != nil {
			return err
		} else {
			if ch, err = StreamCreateFile(p.Conf.Source); err != nil {
				return err
			}
		}
	}
	p.Channel = ch
	if _, err = ChannelSetVolume(p.Channel, p.Conf.ChVol); err != nil {
		log.Println(err)
	}

	if _, err = SetVol(p.Conf.Vol); err != nil {
		return err
	}

	if _, err = ChannelPlay(p.Channel); err != nil {
		return err
	}
	return nil
}

// Player.Pause
func (p *Player) Pause() (err error) {
	if _, err = ChannelPause(p.Channel); err != nil {
		return err
	}
	return nil
}

// Player.Stop
func (p *Player) Stop() (err error) {
	if _, err = ChannelStop(p.Channel); err != nil {
		return err
	}
	return nil
}

// Player.GetChVol
func (p *Player) GetChVol() float32 {
	return p.Conf.ChVol
}

// Player.SetChVol
func (p *Player) SetChVol(v float32) (err error) {
	if _, err = ChannelSetVolume(p.Channel, v); err != nil {
		return err
	}
	p.Conf.ChVol = v
	return nil
}

// Player.GetVol
func (p *Player) GetVol() float32 {
	return p.Conf.Vol
}

// Player.SetVol
func (p *Player) SetVol(v float32) (err error) {
	if _, err = SetVol(v); err != nil {
		return err
	}
	p.Conf.Vol = v
	return nil
}

// Player.Source
func (p *Player) NewSource(src string) {
	p.Conf.Source = src
}

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
	Volume float32 //default 25%
	Source string
}

// NewPlayer
func NewPlayer(conf PlayerConf) (*Player, error) {
	if conf.Device == 0 {
		conf.Device = -1
	}
	if conf.Volume == 0.0 {
		conf.Volume = 25
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
	if _, err = ChannelSetVolume(p.Channel, p.Conf.Volume); err != nil {
		log.Println(err)
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

// Player.GetVol
func (p *Player) GetVol() float32 {
	return p.Conf.Volume
}

// Player.SetVol
func (p *Player) SetVol(v float32) (err error) {
	if _, err = ChannelSetVolume(p.Channel, v); err != nil {
		return err
	}
	p.Conf.Volume = v
	return nil
}

# gobass.dll
Go bindings for [bass.dll] (http://www.un4seen.com/)

Note: wrapped a few basic functions

### Install
```sh
$ go get github.com/lexesv/gobass.dll
```

### Usage

Example 1
```go
if ok, err := bass.Init(-1, 44100, 0); ok {
	fmt.Println("bass init")
} else {
	fmt.Println(err)
}
c, err := bass.StreamCreateURL("http://music.myradio.ua:8000/PopRock_news128.mp3")
if err == nil {
	fmt.Println(bass.ChannelPlay(c))
} else {
	fmt.Println(err)
}
bass.SetVol(50)
bass.ChannelSetVolume(c, 40.5)
```

Example 2
```go
plugin, err := bass.PluginLoad("libbass_aac.so")
	if err != nil {
		fmt.Println(err)
	} else {
		defer bass.PluginFree(plugin)
	}
	
cfg := bass.PlayerConf{
	Device: -1,
	Freq:   44100,
	Flags:  0,
	Volume: 50.5,
	Source: "http://online-hitfm.tavrmedia.ua/HitFM_Live",
}

// or
cfg = bass.PlayerConf{}
cfg.Source = "http://music.myradio.ua:8000/main_stream_rock_news128.mp3"
	
player, err := bass.NewPlayer(cfg)
	
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(player.Play())
	fmt.Println("Volume:", player.GetVol())
	time.Sleep(time.Second * 1)
	player.SetVol(15)
	fmt.Println("Volume:", player.GetVol())
}

r := regexp.MustCompile(`(?isU)StreamTitle='(.*)';`)
for {
	m := r.FindStringSubmatch(bass.ChannelGetTags(player.Channel, bass.BASS_TAG_META))
	if len(m) > 0 {
		fmt.Printf("\r%s", m[1])
	}
	time.Sleep(time.Second * 3)
}
```
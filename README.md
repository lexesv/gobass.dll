# gobass.dll
Go bindings for bass.dll (for minimum needs)

### Install
```sh
$ go get github.com/lexesv/gobass.dll
```

### Usage
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
		Flags:  0, Volume: 50.5,
		Source: "http://online-hitfm.tavrmedia.ua/HitFM_Live",
	}
	
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

for {
	r := regexp.MustCompile(`(?isU)StreamTitle='(.*)';`)
	m := r.FindStringSubmatch(bass.ChannelGetTags(player.Channel, bass.BASS_TAG_META))
	if len(m) > 0 {                                                       ")
		fmt.Printf("\r%s", m[1])
	}
	time.Sleep(time.Second * 3)
}
```
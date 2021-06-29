package xiamicontent

// AudioDto 
type AudioDto struct {

    // 音频ID
    
    AudioId   int64 `json:"audio_id,omitempty" xml:"audio_id,omitempty"`
    

    // 音频时长，同一歌曲时长一致（单位：毫秒）
    
    Duration   int64 `json:"duration,omitempty" xml:"duration,omitempty"`
    

    // 音频地址 带有标准的CDN签名
    
    ListenUrl   string `json:"listen_url,omitempty" xml:"listen_url,omitempty"`
    

    // 文件大小(单位: 字节)
    
    FileSize   int64 `json:"file_size,omitempty" xml:"file_size,omitempty"`
    

    // 音频格式:mp3|m4a|wav
    
    Format   string `json:"format,omitempty" xml:"format,omitempty"`
    

    // 码率: 32|64|128|192|320|>320的整数值
    
    Rate   int64 `json:"rate,omitempty" xml:"rate,omitempty"`
    

    // 品质:1(无损)|2(320k)|3(192k)|4(128k)|5(32k)|7(64k)|15(MQA)
    
    Quality   int64 `json:"quality,omitempty" xml:"quality,omitempty"`
    

    // 采样率
    
    SampleRate   int64 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
    

    // 位信息，音频压缩过程产生的信息，可忽略
    
    Bits   int64 `json:"bits,omitempty" xml:"bits,omitempty"`
    

    // 过期时间(unix timestamp)
    
    Expire   int64 `json:"expire,omitempty" xml:"expire,omitempty"`
    

}

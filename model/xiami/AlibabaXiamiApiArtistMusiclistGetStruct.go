package xiami

// AlibabaXiamiApiArtistMusiclistGetStruct 
/* model for simplify = false
type AlibabaXiamiApiArtistMusiclistGetStruct struct {

    // 艺人信息
    
    Artists  struct {
        AlibabaXiamiApiArtistMusiclistGetStruct  []AlibabaXiamiApiArtistMusiclistGetStruct `json:"alibaba_xiami_api_artist_musiclist_get_struct,omitempty"`
    } `json:"artists,omitempty"`
    

    // 艺人是否有演出的标记
    
    IsShow   bool `json:"is_show,omitempty"`
    

    // 艺人名字
    
    ArtistName   string `json:"artist_name,omitempty"`
    

    // 是否是音乐人
    
    IsMusician   bool `json:"is_musician,omitempty"`
    

    // 粉丝数
    
    CountLikes   int64 `json:"count_likes,omitempty"`
    

    // 艺人id, BIGINT类型
    
    ArtistId   int64 `json:"artist_id,omitempty"`
    

    // 艺人头像
    
    ArtistLogo   string `json:"artist_logo,omitempty"`
    

    // 拼音
    
    Pinyin   string `json:"pinyin,omitempty"`
    

}
*/

// AlibabaXiamiApiArtistMusiclistGetStruct 
type AlibabaXiamiApiArtistMusiclistGetStruct struct {

    // 艺人信息
    Artists   []AlibabaXiamiApiArtistMusiclistGetStruct `json:"artists,omitempty"`

    // 艺人是否有演出的标记
    IsShow   bool `json:"is_show,omitempty"`

    // 艺人名字
    ArtistName   string `json:"artist_name,omitempty"`

    // 是否是音乐人
    IsMusician   bool `json:"is_musician,omitempty"`

    // 粉丝数
    CountLikes   int64 `json:"count_likes,omitempty"`

    // 艺人id, BIGINT类型
    ArtistId   int64 `json:"artist_id,omitempty"`

    // 艺人头像
    ArtistLogo   string `json:"artist_logo,omitempty"`

    // 拼音
    Pinyin   string `json:"pinyin,omitempty"`

}

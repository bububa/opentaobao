package openmall

// TopParseAddressEntryVO 
type TopParseAddressEntryVO struct {

    // 地区编码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 地域名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 地区级别，2代表省、自治区、直辖市、特别行政区；3代表地级市、 地区、盟、自治州；4代表县、区、自治县、旗；5代表乡、镇、街道，openmall中请取第三或者第四级别地域编码传入即可
    
    Scope   int64 `json:"scope,omitempty" xml:"scope,omitempty"`
    

}

package security

// Element 
type Element struct {

    // 名称, 常用名称：姓名(Name); 身份证(Identity_Code); 手机(Mobile); 身份证正面(Identity_Front_Pic);身份证背面(Identity_back_pic);半身照(UpperBody_pic);地址(Address)
    Name   string `json:"name,omitempty"`

    // 值，如果为oss地址，格式如下oss://bucketname:xxxx.jpg
    Value   string `json:"value,omitempty"`

    // 类型, 如果是姓名，身份证，则是ValueType；如果是图片则是PhotoType；如果是JSON格式，则是JsonType
    ValueMeta   string `json:"value_meta,omitempty"`

}

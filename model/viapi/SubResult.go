package viapi

// SubResult 
type SubResult struct {

    // 建议您执行的操作，取值范围： pass：图片正常，无需进行其余操作；或者未识别出目标对象 review：检测结果不确定，需要进行人工审核；或者识别出目标对象 block：图片违规，建议执行进一步操作（如直接删除或做限制处理）
    
    Suggestion   string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
    

    // 相似概率
    
    Rate   int64 `json:"rate,omitempty" xml:"rate,omitempty"`
    

    // 检测结果的分类，与具体的scene对应。取值范围参见scene和label说明。
    
    Label   string `json:"label,omitempty" xml:"label,omitempty"`
    

    // 图片检测场景，和调用请求中的场景（scenes）对应
    
    Scene   string `json:"scene,omitempty" xml:"scene,omitempty"`
    

    // 识别到的图片中的完整文字信息。 说明 默认不返回，如需返回请通过工单联系我们
    
    OcrDataList   []string `json:"ocr_data_list,omitempty" xml:"ocr_data_list>string,omitempty"`
    

    // 如果待检测图片因为过长被截断，该参数返回截断后的每一帧图像的临时访问地址，供您参考。具体结构描述见Frame
    
    Frames   []Frame `json:"frames,omitempty" xml:"frames,omitempty"`
    

    // 图片中含有广告时，返回图片中广告文字命中的风险关键词信息。格式为数组，具体结构描述见HintWordsInfo。 说明 仅适用于ad场景。 示例值： "hintWordsInfo":[{"context":"敏感词"}]
    
    HintWordsInfoList   []HintWordsInfo `json:"hint_words_info_list,omitempty" xml:"hint_words_info_list,omitempty"`
    

    // 图片中含有小程序码时，返回小程序码的位置信息，具体结构描述见ProgramCodeData。说明 仅适用于qrcode场景，且已通过工单联系我们开通了小程序码识别
    
    ProgramCodeDataList   []ProgramCodeData `json:"program_code_data_list,omitempty" xml:"program_code_data_list,omitempty"`
    

    // 图片中含有logo时，返回识别出来的logo信息，具体结构描述见LogoData。 说明 仅适用于logo场景
    
    LogoDataList   []LogoData `json:"logo_data_list,omitempty" xml:"logo_data_list,omitempty"`
    

    // 图片中包含暴恐识涉政内容时，返回识别出来的暴恐涉政信息，具体结构描述见SfaceData。 说明 仅适用于terrorism和sface场景。关于该参数在sface场景中的具体内容，请参见敏感人脸检测
    
    SfaceDataList   []SfaceData `json:"sface_data_list,omitempty" xml:"sface_data_list,omitempty"`
    

}

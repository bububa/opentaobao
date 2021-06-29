package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取手机banner图 APIResponse
alibaba.xiami.api.mobile.figureimage.get

获取手机banner图
*/
type AlibabaXiamiApiMobileFigureimageGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiMobileFigureimageGetResponse
}

type AlibabaXiamiApiMobileFigureimageGetResponse struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_mobile_figureimage_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // mobile_figure_image_result
    
    MobileFigureImageList   []AlibabaXiamiApiMobileFigureimageGetStruct `json:"mobile_figure_image_list,omitempty" xml:"mobile_figure_image_list>alibaba_xiami_api_mobile_figureimage_get_struct,omitempty"`
    
    
}

package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取手机banner图 API返回值 
alibaba.xiami.api.mobile.figureimage.get

获取手机banner图
*/
type AlibabaXiamiApiMobileFigureimageGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiMobileFigureimageGetAPIResponseModel
}

// 获取手机banner图 成功返回结果
type AlibabaXiamiApiMobileFigureimageGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_mobile_figureimage_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // mobile_figure_image_result
    MobileFigureImageList   []AlibabaXiamiApiMobileFigureimageGetStruct `json:"mobile_figure_image_list,omitempty" xml:"mobile_figure_image_list>alibaba_xiami_api_mobile_figureimage_get_struct,omitempty"`
}

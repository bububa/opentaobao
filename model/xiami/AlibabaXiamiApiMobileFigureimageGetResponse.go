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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_mobile_figureimage_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // mobile_figure_image_result
    
    MobileFigureImageList   []AlibabaXiamiApiMobileFigureimageGetStruct `json:"mobile_figure_image_list,omitempty" xml:"
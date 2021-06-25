package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取手机banner图 APIResponse
alibaba.xiami.api.mobile.figureimage.get

获取手机banner图
*/
type AlibabaXiamiApiMobileFigureimageGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiMobileFigureimageGetResponse `json:"alibaba_xiami_api_mobile_figureimage_get_response,omitempty"`
}

type AlibabaXiamiApiMobileFigureimageGetResponse struct {

    // mobile_figure_image_result
    MobileFigureImageList   []AlibabaXiamiApiMobileFigureimageGetStruct `json:"mobile_figure_image_list,omitempty"`

}

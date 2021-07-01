package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上线/下线 体检产品 API返回值 
alibaba.alihealth.examination.goods.onoffline

第三方体检机构对接钉钉体检中的产品 上线／下线
*/
type AlibabaAlihealthExaminationGoodsOnofflineAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationGoodsOnofflineAPIResponseModel
}

// 上线/下线 体检产品 成功返回结果
type AlibabaAlihealthExaminationGoodsOnofflineAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_goods_onoffline_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

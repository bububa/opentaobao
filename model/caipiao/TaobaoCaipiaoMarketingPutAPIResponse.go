package caipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoMarketingPutAPIResponse 创建或修改商家送彩票活动 API返回值
// taobao.caipiao.marketing.put
//
// 卖家通过此接口新增或修改送彩票活动的配置，比如活动时间、活动的条件等。
//
// 店铺装修、宝贝详情页的宣导素材示例：
// https://gw.alicdn.com/tfs/TB1_nOiSXXXXXbgXXXXXXXXXXXX-750-280.png
// https://gw.alicdn.com/tfs/TB1FZX6SXXXXXXzXFXXXXXXXXXX-790-280.png
// https://gw.alicdn.com/tfs/TB1z4t8SXXXXXckXpXXXXXXXXXX-750-280.png
// https://gw.alicdn.com/tfs/TB1BhqgSXXXXXcDXXXXXXXXXXXX-750-280.png
// https://gw.alicdn.com/tfs/TB1TYt9SXXXXXXAXFXXXXXXXXXX-750-280.png
// https://gw.alicdn.com/tfs/TB1tzpNSXXXXXacXVXXXXXXXXXX-790-280.png
// https://gw.alicdn.com/tfs/TB1UXdxSXXXXXXsapXXXXXXXXXX-790-280.png
// https://gw.alicdn.com/tfs/TB1_gV.SXXXXXbZXpXXXXXXXXXX-790-280.png
type TaobaoCaipiaoMarketingPutAPIResponse struct {
	model.CommonResponse
	TaobaoCaipiaoMarketingPutAPIResponseModel
}

// TaobaoCaipiaoMarketingPutAPIResponseModel is 创建或修改商家送彩票活动 成功返回结果
type TaobaoCaipiaoMarketingPutAPIResponseModel struct {
	XMLName xml.Name `xml:"caipiao_marketing_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务操作结果,true成功/false失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

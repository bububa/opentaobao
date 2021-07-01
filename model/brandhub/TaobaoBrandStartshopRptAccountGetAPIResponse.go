package brandhub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
明星店铺账户报表数据查询 API返回值 
taobao.brand.startshop.rpt.account.get

获取明星店铺广告主账户整体报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
type TaobaoBrandStartshopRptAccountGetAPIResponse struct {
    model.CommonResponse
    TaobaoBrandStartshopRptAccountGetAPIResponseModel
}

// 明星店铺账户报表数据查询 成功返回结果
type TaobaoBrandStartshopRptAccountGetAPIResponseModel struct {
    XMLName xml.Name `xml:"brand_startshop_rpt_account_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    ErrorParam   string `json:"error_param,omitempty" xml:"error_param,omitempty"`
    // result
    AdvertiserRptList   []TaobaoBrandStartshopRptAccountGetResult `json:"advertiser_rpt_list,omitempty" xml:"advertiser_rpt_list>taobao_brand_startshop_rpt_account_get_result,omitempty"`
}

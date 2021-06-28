package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户售后服务模板 APIResponse
taobao.aftersale.get

查询用户设置的售后服务模板，仅返回标题和id
*/
type TaobaoAftersaleGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"aftersale_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 售后服务返回对象
    
    AfterSales   []AfterSale `json:"after_sales,omitempty" xml:"
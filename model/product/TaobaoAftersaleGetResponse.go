package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询用户售后服务模板 APIResponse
taobao.aftersale.get

查询用户设置的售后服务模板，仅返回标题和id
*/
type TaobaoAftersaleGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAftersaleGetResponse `json:"taobao_aftersale_get_response,omitempty"`
}

type TaobaoAftersaleGetResponse struct {

    // 售后服务返回对象
    AfterSales   []AfterSale `json:"after_sales,omitempty"`

}

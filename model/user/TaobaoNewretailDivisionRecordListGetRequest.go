package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导购分佣明细列表 API请求
taobao.newretail.division.record.list.get

提供分页查询导购分佣明细的能力
*/
type TaobaoNewretailDivisionRecordListGetAPIRequest struct {
    model.Params
    // 入参
    _param   *TopDivisionRecordReqDTO
}

// 初始化TaobaoNewretailDivisionRecordListGetAPIRequest对象
func NewTaobaoNewretailDivisionRecordListGetRequest() *TaobaoNewretailDivisionRecordListGetAPIRequest{
    return &TaobaoNewretailDivisionRecordListGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoNewretailDivisionRecordListGetAPIRequest) GetApiMethodName() string {
    return "taobao.newretail.division.record.list.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoNewretailDivisionRecordListGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *TaobaoNewretailDivisionRecordListGetAPIRequest) SetParam(_param *TopDivisionRecordReqDTO) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoNewretailDivisionRecordListGetAPIRequest) GetParam() *TopDivisionRecordReqDTO {
    return r._param
}

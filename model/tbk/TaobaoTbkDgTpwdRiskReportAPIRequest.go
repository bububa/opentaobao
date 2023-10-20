package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgtpwdriskreportAPIRequest 淘宝客-推广者-淘口令预警及拦截查询 API请求
// taobao.tbk.dg.tpwd.risk.report
//
// 淘宝客-推广者-淘口令预警及拦截查询
type TaobaotbkdgtpwdriskreportAPIRequest struct {
	model.Params
	// 如有pid，则查询pid下的relationid名单；如没有pid，则查询appkey关联userid下的pid名单
	_pid string
	// 分页参数
	_offset int64
	// 分页参数，一次最多不能超过1000
	_limit int64
}

// NewTaobaotbkdgtpwdriskreportRequest 初始化TaobaotbkdgtpwdriskreportAPIRequest对象
func NewTaobaotbkdgtpwdriskreportRequest() *TaobaotbkdgtpwdriskreportAPIRequest {
	return &TaobaotbkdgtpwdriskreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgtpwdriskreportAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.tpwd.risk.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgtpwdriskreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgtpwdriskreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPid is Pid Setter
// 如有pid，则查询pid下的relationid名单；如没有pid，则查询appkey关联userid下的pid名单
func (r *TaobaotbkdgtpwdriskreportAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaotbkdgtpwdriskreportAPIRequest) GetPid() string {
	return r._pid
}

// SetOffset is Offset Setter
// 分页参数
func (r *TaobaotbkdgtpwdriskreportAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaotbkdgtpwdriskreportAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetLimit is Limit Setter
// 分页参数，一次最多不能超过1000
func (r *TaobaotbkdgtpwdriskreportAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r TaobaotbkdgtpwdriskreportAPIRequest) GetLimit() int64 {
	return r._limit
}

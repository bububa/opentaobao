package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscmembergroupoptionalAPIRequest 工具服务商member组查询、新增接口 API请求
// taobao.tbk.sc.membergroup.optional
//
// 工具服务商member组查询、新增接口
type TaobaotbkscmembergroupoptionalAPIRequest struct {
	model.Params
	// 淘宝数字id
	_tbnumids string
	// member组id
	_membergroupid int64
}

// NewTaobaotbkscmembergroupoptionalRequest 初始化TaobaotbkscmembergroupoptionalAPIRequest对象
func NewTaobaotbkscmembergroupoptionalRequest() *TaobaotbkscmembergroupoptionalAPIRequest {
	return &TaobaotbkscmembergroupoptionalAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscmembergroupoptionalAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.membergroup.optional"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscmembergroupoptionalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscmembergroupoptionalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbnumids is Tbnumids Setter
// 淘宝数字id
func (r *TaobaotbkscmembergroupoptionalAPIRequest) SetTbnumids(_tbnumids string) error {
	r._tbnumids = _tbnumids
	r.Set("tb_num_ids", _tbnumids)
	return nil
}

// GetTbnumids Tbnumids Getter
func (r TaobaotbkscmembergroupoptionalAPIRequest) GetTbnumids() string {
	return r._tbnumids
}

// SetMembergroupid is Membergroupid Setter
// member组id
func (r *TaobaotbkscmembergroupoptionalAPIRequest) SetMembergroupid(_membergroupid int64) error {
	r._membergroupid = _membergroupid
	r.Set("member_group_id", _membergroupid)
	return nil
}

// GetMembergroupid Membergroupid Getter
func (r TaobaotbkscmembergroupoptionalAPIRequest) GetMembergroupid() int64 {
	return r._membergroupid
}

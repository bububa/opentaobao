package pentraprism

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopentaprismtaskqueryitemAPIRequest 查询任务当前进度 API请求
// taobao.pentaprism.task.queryitem
//
// 外网用户查询五棱镜任务系统当前进度
type TaobaopentaprismtaskqueryitemAPIRequest struct {
	model.Params
	// TOP接口标准入参
	_openPo *OpenTaskPo
}

// NewTaobaopentaprismtaskqueryitemRequest 初始化TaobaopentaprismtaskqueryitemAPIRequest对象
func NewTaobaopentaprismtaskqueryitemRequest() *TaobaopentaprismtaskqueryitemAPIRequest {
	return &TaobaopentaprismtaskqueryitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopentaprismtaskqueryitemAPIRequest) GetApiMethodName() string {
	return "taobao.pentaprism.task.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopentaprismtaskqueryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopentaprismtaskqueryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenPo is OpenPo Setter
// TOP接口标准入参
func (r *TaobaopentaprismtaskqueryitemAPIRequest) SetOpenPo(_openPo *OpenTaskPo) error {
	r._openPo = _openPo
	r.Set("open_po", _openPo)
	return nil
}

// GetOpenPo OpenPo Getter
func (r TaobaopentaprismtaskqueryitemAPIRequest) GetOpenPo() *OpenTaskPo {
	return r._openPo
}

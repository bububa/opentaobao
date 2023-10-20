package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminAdmOttAuditAPIRequest 优酷OTT广告素材审核 API请求
// yunos.tvpubadmin.adm.ott.audit
//
// 审核优酷OTT端广告素材
type YunosTvpubadminAdmOttAuditAPIRequest struct {
	model.Params
	// 广告审核内容，json格式
	_data string
}

// NewYunosTvpubadminAdmOttAuditRequest 初始化YunosTvpubadminAdmOttAuditAPIRequest对象
func NewYunosTvpubadminAdmOttAuditRequest() *YunosTvpubadminAdmOttAuditAPIRequest {
	return &YunosTvpubadminAdmOttAuditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminAdmOttAuditAPIRequest) Reset() {
	r._data = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminAdmOttAuditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.adm.ott.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminAdmOttAuditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminAdmOttAuditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 广告审核内容，json格式
func (r *YunosTvpubadminAdmOttAuditAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r YunosTvpubadminAdmOttAuditAPIRequest) GetData() string {
	return r._data
}

var poolYunosTvpubadminAdmOttAuditAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminAdmOttAuditRequest()
	},
}

// GetYunosTvpubadminAdmOttAuditRequest 从 sync.Pool 获取 YunosTvpubadminAdmOttAuditAPIRequest
func GetYunosTvpubadminAdmOttAuditAPIRequest() *YunosTvpubadminAdmOttAuditAPIRequest {
	return poolYunosTvpubadminAdmOttAuditAPIRequest.Get().(*YunosTvpubadminAdmOttAuditAPIRequest)
}

// ReleaseYunosTvpubadminAdmOttAuditAPIRequest 将 YunosTvpubadminAdmOttAuditAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminAdmOttAuditAPIRequest(v *YunosTvpubadminAdmOttAuditAPIRequest) {
	v.Reset()
	poolYunosTvpubadminAdmOttAuditAPIRequest.Put(v)
}

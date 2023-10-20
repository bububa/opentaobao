package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseadminthemehotupdate 云主题热更新数据集
// alibaba.alihouse.admin.theme.hot.update
//
// 云主题更新
func Alibabaalihouseadminthemehotupdate(clt *core.SDKClient, req *alihouse.AlibabaalihouseadminthemehotupdateAPIRequest, session string) (*alihouse.AlibabaalihouseadminthemehotupdateAPIResponse, error) {
	var resp alihouse.AlibabaalihouseadminthemehotupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

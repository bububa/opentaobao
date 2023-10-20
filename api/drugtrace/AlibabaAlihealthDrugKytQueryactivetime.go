package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytqueryactivetime 药品激活状态同步
// alibaba.alihealth.drug.kyt.queryactivetime
//
// 根据赋码资源（CodeVersion + resCode）获得最新激活时间
// 应用于各地市对接前进行药品目录匹配，医保中心存在的药品可能比较陈旧杂乱
func Alibabaalihealthdrugkytqueryactivetime(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytqueryactivetimeAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytqueryactivetimeAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytqueryactivetimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

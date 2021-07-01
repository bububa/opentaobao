package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaCfdaXtptAppAcceptInfo
协同平台数据下行接口
alibaba.cfda.xtpt.app.accept.info

协同平台数据下行接口 */
func AlibabaCfdaXtptAppAcceptInfo(clt *core.SDKClient, req *drugtrace.AlibabaCfdaXtptAppAcceptInfoAPIRequest, session string) (*drugtrace.AlibabaCfdaXtptAppAcceptInfoAPIResponse, error) {
	var resp drugtrace.AlibabaCfdaXtptAppAcceptInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

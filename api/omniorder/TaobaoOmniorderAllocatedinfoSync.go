package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderallocatedinfosync 分单结果同步给星盘
// taobao.omniorder.allocatedinfo.sync
//
// ISV分单完成，将分单结果同步给星盘
func Taobaoomniorderallocatedinfosync(clt *core.SDKClient, req *omniorder.TaobaoomniorderallocatedinfosyncAPIRequest, session string) (*omniorder.TaobaoomniorderallocatedinfosyncAPIResponse, error) {
	var resp omniorder.TaobaoomniorderallocatedinfosyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

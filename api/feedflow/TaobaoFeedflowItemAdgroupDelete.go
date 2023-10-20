package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgroupdelete 根据单元id删除单元
// taobao.feedflow.item.adgroup.delete
//
// 根据单元id删除单元
func Taobaofeedflowitemadgroupdelete(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgroupdeleteAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgroupdeleteAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgroupdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

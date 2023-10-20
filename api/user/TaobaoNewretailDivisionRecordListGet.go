package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Taobaonewretaildivisionrecordlistget 导购分佣明细列表
// taobao.newretail.division.record.list.get
//
// 提供分页查询导购分佣明细的能力
func Taobaonewretaildivisionrecordlistget(clt *core.SDKClient, req *user.TaobaonewretaildivisionrecordlistgetAPIRequest, session string) (*user.TaobaonewretaildivisionrecordlistgetAPIResponse, error) {
	var resp user.TaobaonewretaildivisionrecordlistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

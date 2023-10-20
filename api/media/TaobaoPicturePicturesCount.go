package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// Taobaopicturepicturescount 图片总数查询
// taobao.picture.pictures.count
//
// 图片总数查询，目前出于对数据库的保护暂不支持此功能
func Taobaopicturepicturescount(clt *core.SDKClient, req *media.TaobaopicturepicturescountAPIRequest, session string) (*media.TaobaopicturepicturescountAPIResponse, error) {
	var resp media.TaobaopicturepicturescountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

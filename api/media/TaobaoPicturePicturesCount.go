package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPicturePicturesCount 图片总数查询
// taobao.picture.pictures.count
//
// 图片总数查询，目前出于对数据库的保护暂不支持此功能
func TaobaoPicturePicturesCount(clt *core.SDKClient, req *media.TaobaoPicturePicturesCountAPIRequest, resp *media.TaobaoPicturePicturesCountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}

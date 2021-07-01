package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPicturePicturesCountAPIRequest
图片总数查询 API请求
taobao.picture.pictures.count

图片总数查询 */
type TaobaoPicturePicturesCountAPIRequest struct {
	model.Params
	// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
	_startDate string
	// 图片分类
	_pictureCategoryId int64
	// 文件名
	_title string
	// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
	_endDate string
	// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
	_startModifiedDate string
	// 是否删除,undeleted代表没有删除,deleted表示删除
	_deleted string
	// 图片ID
	_pictureId int64
	// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
	_clientType string
}

// New

package model

type Product struct {
	BaseModel
	Category
	Brands
	BrandsId    int32    `gorm:"type:int;not null"`
	CategoryId  int32    `gorm:"type:int;not null"`
	OnSales     bool     `gorm:"type:bool;default:false;not null" json:"onSales"`  //是否上架
	ShipFree    bool     `gorm:"type:bool;default:false;not null" json:"shipFree"` //是否免运费
	IsNew       bool     `gorm:"type:bool;default:false; not null" json:"isNew"`   //是否是新品
	IsHot       bool     `gorm:"type:bool;default:false;not null" json:"isHot"`    //是否是热销商品
	Name        string   `gorm:"type:varchar(64);not null" json:"name"`            //商品名称
	ProductSn   string   `gorm:"type:varchar(64);not null" json:"productSn"`       // 商品编号
	ClickNum    int32    `gorm:"type:int;default:1;not null" json:"clickNum"`      //点击数
	SoldNum     int32    `gorm:"type:int; default:0;not null" json:"soldNum"`      // 商品销售数
	FavNum      int32    `gorm:"type:int;default:0;not null" json:"favNum"`        // 商品收藏数量
	MarketPrice float32  `gorm:"type:int;not null" json:"marketPrice"`             // 商吕标签价
	ShopPrice   float32  `gorm:"type:int;not null" json:"shopPrice"`               //商品销售价
	Description string   `gorm:"type:varchar(200);not null" json:"productBrief"`   // 商品简介
	Images      GormList `gorm:"type:varchar(1000);not null" json:"images"`        // 商品图片
	DescImages  GormList `gorm:"type:varchar(100);not null" json:"descImages"`     // 商品详情图片
	CoverImage  string   `gorm:"type:varchar(200);not null" json:"CoverImage"`     //商品封面图
}

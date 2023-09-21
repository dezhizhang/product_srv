package model

type Banner struct {
	BaseModel
	Name        string `gorm:"type:varchar(32);" json:"name"`
	Status      string `gorm:"type:varchar(32);not null" json:"status"`
	Position    string `gorm:"type:varchar(32);not nul" json:"position"`
	Link        string `gorm:"type:varchar(200);not null" json:"link"`
	Url         string `gorm:"type:varchar(200);not null" json:"url"`
	Description string `gorm:"type:varchar(200)" json:"description"`
}

type ElasticBanner struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (ElasticBanner) GetIndexName() string {
	return "banner"
}

func (ElasticBanner) GetMapping() string {
	mapping := `
		"mappings":{
			"properties":{
				"id":{
					"type":"integer"
				},
				"name":{
					"type":"text",
					"analyzer":"ik_max_word"
					"fields":{
						"keyword":{
							"type":"keyword",
							"ignore_above":256
						}
					}
				}
				
			}
		}	
	`
	return mapping
}

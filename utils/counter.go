package utils

type NameCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Counter struct {
	Count int         `json:"count"`
	Data  []NameCount `json:"data"`
}

// 全体カウントを +1
func (c *Counter) IncrementTotal() {
	c.Count++
}

// 名前付きカウントを +1（なければ追加）
func (c *Counter) IncrementByName(name string) {
	for i, item := range c.Data {
		if item.Name == name {
			c.Data[i].Count++
			return
		}
	}
	// 見つからなければ新規追加
	c.Data = append(c.Data, NameCount{Name: name, Count: 1})
}
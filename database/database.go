package database

type dbItem struct {
	value interface{}
}

type DB struct {
	num   int
	items map[string]dbItem
}

func NewDB(n int) *DB {
	return &DB{
		num:   n,
		items: make(map[string]dbItem),
	}
}

func (db *DB) Set(key string, value interface{}) {
	db.items[key] = dbItem{value: value}
}

func (db *DB) Get(key string) (interface{}, bool) {
	item, ok := db.items[key]
	if !ok {
		return nil, false
	}
	return item.value, true
}

package products


	productH := &Handler{
		store: &datastore{
			m: map[string]products.Product{
				"1": {ID: "1", Name: "Tomates", Image: "tomates.png", Total: 1, Price: 5.00, SoldOut: true},
			},
			RWMutex: &sync.RWMutex{},
		},
	}

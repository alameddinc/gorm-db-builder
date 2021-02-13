
# gorm-db-builder

**Türkçe Kullanım Örneği**
Detaylı Kullanım örneklerini aşağıdaki makaleden inceleyebilirsiniz.
https://alameddinc.medium.com/gorm-db-için-arayüz-geliştirmek-4df6fd641840


## Örnek Kullanımlar (V1)

**FetchOne** 
```
func (p *Class) FetchOne(with ...string) error {  
   return DB.FetchOne(p, nil, with...)  
}
```

**FetchAll**
```
type Cities []City  
  
func (p *Cities) FetchAll(where *City, with ...string) error {  
   if where == nil {  
      where = &City{}  
   }  
   return DB.FetchAll(p, where, nil, with...)  
}
```

**Update**
```
func (p *City) Update(update *City) error {  
   return DB.Update(p, update, nil)  
}

//Bulk Update
type Cities []City 

func (p *Cities) UpdateBulk(update *City) error {  
   return DB.Update(p, update, nil)  
}
```

**Save**
```
func (p *City) Save() error {  
   err := DB.Save(p)  
   if err != nil {  
      return err  
   }  
   return p.FetchOne()  
}

// Bulk Save
type Cities []City  
  
func (p *Cities) SaveBulk() error {  
   err := DB.Save(p)  
   if err != nil {  
      return err  
   }  
   return p.FetchOne()  
}
```

**FetchAll**
```

type Cities []City  
func (p *Cities) DeleteBulk() error {  
   return DB.Remove(p, nil)  
}
```

**Remove**
```
func (p *City) Delete() error {  
   return DB.Remove(p, nil)  
}

//Bulk Remove
type Cities []City  
  
func (p *Cities) FetchAll(where *City, with ...string) error {  
   return DB.Remove(p, nil)  
}
```

**FetchOneWithId**
```
func (p *City) FetchOneWithId(id int, with ...string) error {  
   return DB.FetchOneWithID(p, id, nil, with...) 
}
```

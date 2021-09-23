# go-api
A easy way to implements RESTful API views using Fiber and GORM

## Implementation:

Set the enviroment variable to indicates the database:
DB_DNS="postgresql://go_api:go_api@localhost:5432/go_api"


```
# views.go

import "github.com/ikaromn/go-api"


var myModel MyModel
var myModels []MyModel


var MyView = View{
    model:       myModel,
    models:      myModels,
    lookupField: "id",
}

```

```
# routes.go

app.Get("/my-view/:id?", MyView.GetAndListView)
app.Post("/my-view/", MyView.CreateView)
app.Patch("/my-view/", MyView.PartialUpdateView)

```

# Prueba 1

Se consulta la url  a ```https://apirecruit-gjvkhl2c6a-
uc.a.run.app/compras/ ``` desde una fecha especifica y hasta los dias establecidos posterior mente


## BreakPoint

``` html
GET :8080/resumen/2019-12-01?dias=5
```

## Respuesta


``` json
{
"total": 18126265.830000002,
"comprasPorTDC": {
"oro": 18125685,
"amex": 1150
},
"noCompraron": 1177,
"compraMasAlta": 30470.45
}
```
# Prueba 2
Lee el archivo de nombre ```archivo.csv``` 
## uso

desde terminal o el ejecutable desde manejador de archivos
## Respuesta
``` json
[
  {
    "organizacion": "organizacion",
    "users": [
      {
        "username": " usuario",
        "roles": [
          " rol "
        ]
      }
    ]
  },
  {
    "organizacion": "org1",
    "users": [
      {
        "username": "jperez",
        "roles": [
          "admin ",
          "superadmin"
        ]
      },
      {
        "username": "asosa",
        "roles": [
          "writer"
        ]
      }
    ]
  },
  {
    "organizacion": "org2",
    "users": [
      {
        "username": "jperez",
        "roles": [
          "admin "
        ]
      },
      {
        "username": "rrodriguez",
        "roles": [
          "writer",
          "editor"
        ]
      }
    ]
  }
]
```
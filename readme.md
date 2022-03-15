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

## uso
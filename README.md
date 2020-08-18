# Floyd_Go

Encontrar números duplicados em uma sequência utilizando Golang


Problema: Dado uma série de números inteiros mostre pelo menos 1 número que esteja duplicado. Considerando que poderá haver mais de 1 número duplicado e que ele pode duplicar mais de uma vez.
Foi lançado o desafio, para encontrar qual é o número duplicado, e se há mais de um.
Para solucionar esse problema é necessário aplicar no algoritmo a Detecção de ciclo de Floyd, "A Lebre e a Tartaruga".

<img src="https://2.bp.blogspot.com/-Go_uBc1t_zE/XNagA-RKOlI/AAAAAAAADrU/PsygL-cviSckcRSRQhU01R-Cx4zf_RfxACLcBGAs/s1600/floyds-circle.png" align="middle">

A grosso modo intende-se que a lebre é muito mais rápida que a tartaruga, seguindo desse pressuposto, a cada etapa (Y) que a tartaruga alcança no clico de (X), a lebre faz um clico completo (Z). Com isso é possível identificar no clico (Z) se a etapa (Y) se repete alguma vez. Considerando que o ciclo (Z) é formado por todas as etapas(Y) igual ao ciclo(X).
Primeiro definimos qual será a sequencia de números inteiros
slice := []int{9, 3, 9, 3, 6, 3, 6, 7}
Em seguida para cada número apresentado na sequência devemos, devemos fazer ciclo completo em todos os números a frente do atual. 
Obs: esse algoritmo que desenvolvi tem o cuidado de seguir sempre em frente e não verifica o mesmo número mais de uma vez, por exemplo: caso ele encontre um número repetido, ele salva sua posição. Quando é a vez de conferir essa posição para verificar se o número é repetido, ele ignora e pula para o seguinte, pois já sabia que o número em questão já está duplicado.


---
```
for i := 0; i < len(slice); i++ {
dubi = slice[i]
repetiu = false
duplicado = 0
```
### Verifica se a atual posição já foi identificada como duplicada, caso seja pula para o próximo número.

```
if len(posj) > 0 {
for k := 0; k < len(posj); k++ {
if i == posj[k] {
repetiu = true
}
```
### Então confere cada posição com todos o números seguintes da sequência.

```
for j := i + 1; j < len(slice); j++ {
dubj = slice[j]
if dubj == dubi {
duplicado = dubj
dubx++
posj = append(posj, j)
}}
```
### Caso encontre algum número duplicado adiciona ele em um slice

```
if duplicado != 0 {
posi = append(posi, i)
qtdnum = append(qtdnum, duplicado, i, dubx)
repeat := false
if len(slcdubNum) > 0 {
for w := 0; w < len(slcdubNum); w++ {
if duplicado == slcdubNum[w] {
repeat = true
}}
if !repeat {
slcdubNum = append(slcdubNum, duplicado)
}
} else {
slcdubNum = append(slcdubNum, duplicado)
}}
```

[Linkedin](https://www.linkedin.com/in/marcos-rocha-8b3ba158/)
[Medium](https://medium.com/@droidg4)

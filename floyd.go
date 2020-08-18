package main

import (
	"fmt"
)

// encontrar o número duplicado
// dado uma série de numeros inteiros mostre pelo menos 1 número que deve estar duplicado [Pingeonhole principle]
// considerando que há apenas 1 numero duplicado e que ele pode duplicar mais de uma vez
// está lançado o desafio, encontre qual é o número duplicado.
// Para solucionar esse problema é necessário aplicar no algoritmo do pricipio de Floyd's Tortoise and Hare
func main() { // 0, 1, 2, 3, 4, 5 ,6, 7
	slice := []int{9, 3, 9, 3, 6, 3, 6, 7}
	// as váriáveis #dubi & #dubj foram criadas para identificar com facilidade
	// qual o laço de repetição ela faz parte
	dubi := 0
	dubj := 0
	duplicado := 0
	dubx := 0
	repetiu := false
	slcdubNum := []int{} //armazenamento de números duplicados
	slcqtdpos := []int{}
	posi := []int{}
	posj := []int{}
	nums := [][]int{}
	qtdnum := []int{}

	// Percorrer todo o slice
	for i := 0; i < len(slice); i++ {
		dubi = slice[i]
		repetiu = false
		duplicado = 0

		//verifica se a atual posição já foi identificada como duplicada
		// caso já seja uma posição duplicada, pula para o próximo número
		if len(posj) > 0 {
			for k := 0; k < len(posj); k++ {
				if i == posj[k] {
					repetiu = true
				}
			}
			if repetiu {
				continue
			}
		}
		if dubx != 0 {
			slcqtdpos = append(slcqtdpos, dubx)
		}
		dubx = 0
		//confere cada posição com com todos o numeros da sequencia
		for j := i + 1; j < len(slice); j++ {
			dubj = slice[j]
			if dubj == dubi {
				duplicado = dubj
				dubx++
				posj = append(posj, j)
			}
		}
		//Caso encontre algum número duplicado adiciona ele no slice
		if duplicado != 0 {
			posi = append(posi, i)
			qtdnum = append(qtdnum, duplicado, i, dubx)
			repeat := false
			if len(slcdubNum) > 0 {
				for w := 0; w < len(slcdubNum); w++ {
					if duplicado == slcdubNum[w] {
						repeat = true
					}
				}
				if !repeat {
					slcdubNum = append(slcdubNum, duplicado)
				}
			} else {
				slcdubNum = append(slcdubNum, duplicado)
			}
		}
	}

	nums = append(nums, slcdubNum, posi, slcqtdpos, posj)
	fmt.Println("Números duplicados:", slcdubNum)
	fmt.Println("Posições do Número:", posi)
	fmt.Println("Quantidade de vezes que o número foi duplicado:", slcqtdpos)
	fmt.Println("Posições dos Números duplicados:", posj)
	fmt.Println("Agrupados:", nums)

}

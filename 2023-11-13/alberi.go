package main

import (
	"fmt"
)

type Dipendente struct {
	nome        string
	subordinati []*Dipendente
	supervisore *Dipendente
}

type Gerarchia *Dipendente

var elenco []*Dipendente

func nuovoDipendente(nome string, subordinati []*Dipendente, supervisore *Dipendente) *Dipendente {
	return &Dipendente{nome, subordinati, supervisore}
}
func nuovaGerarchia(nome string) Gerarchia {
	var sl []*Dipendente
	capo := nuovoDipendente(nome, sl, nil)
	elenco = append(elenco, capo)
	return capo
}

func aggiungiDipendente(nome string, subordinati []*Dipendente, supervisore *Dipendente) *Dipendente {
	dipendente := nuovoDipendente(nome, subordinati, supervisore)
	for {
		supervisore.subordinati = append(supervisore.subordinati, dipendente)
		if supervisore.supervisore != nil {
			supervisore = supervisore.supervisore
		} else {
			break
		}
	}
	elenco = append(elenco, dipendente)
	return dipendente
}

func stampaSubordinati(dipendente *Dipendente) {
	fmt.Print("subordinati di ", dipendente.nome, ": ")
	numSubordinati := len(dipendente.subordinati)
	for i := 0; i < numSubordinati-1; i++ {
		fmt.Print(dipendente.subordinati[i].nome, ", ")
	}
	if numSubordinati != 0 {
		fmt.Println(dipendente.subordinati[numSubordinati-1].nome)
	} else {
		fmt.Println()
	}
}

func senzaSubordinati() []string {
	var noSub []string
	for i := 0; i < len(elenco); i++ {
		dipendente := *elenco[i]
		if len(dipendente.subordinati) == 0 {
			noSub = append(noSub, dipendente.nome)
		}
	}
	return noSub
}

func trovaSupervisore(dipendente *Dipendente) string {
	if dipendente.supervisore != nil {
		return dipendente.supervisore.nome
	} else {
		return "nessuno, è dipendente di massimo livello"
	}
}

func stampaSupervisori(dipendente *Dipendente) {
	fmt.Print("dipendenti di livello superiore di ", dipendente.nome, ": ")
	super := dipendente.supervisore
	for {
		if super != nil {
			fmt.Print(super.nome, " ")
			super = super.supervisore
		} else {
			break
		}
	}
	fmt.Println()
}

func stampaElenco() {
	fmt.Print("elenco dei dipendenti, con eventuali supervisori: ")
	for _, k := range elenco {
		fmt.Print("[", k.nome)
		if k.supervisore != nil {
			fmt.Print(", con supervisore: ", k.supervisore.nome)
		}
		fmt.Print("] ")
	}
	fmt.Println()
}

func main() {
	Anna := nuovaGerarchia("Anna")
	Bruno := aggiungiDipendente("Bruno", nil, Anna)
	Carlo := aggiungiDipendente("Carlo", nil, Anna)
	Daniela := aggiungiDipendente("Daniela", nil, Anna)
	Enrico := aggiungiDipendente("Enrico", nil, Bruno)
	Francesco := aggiungiDipendente("Francesco", nil, Bruno)
	stampaSubordinati(Anna)
	stampaSubordinati(Bruno)
	stampaSubordinati(Carlo)
	stampaSubordinati(Daniela)
	stampaSubordinati(Enrico)
	stampaSubordinati(Francesco)
	fmt.Println("dipendenti che non hanno subordinati: ", senzaSubordinati())
	fmt.Println("supervisore di anna: ", trovaSupervisore(Anna))
	fmt.Println("supervisore di enrico: ", trovaSupervisore(Enrico))
	stampaSupervisori(Francesco)
	stampaElenco()
}

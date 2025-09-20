package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

// Tabela de permutação inicial (IP)
var initialPermutation = []int{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

// Tabela de permutação final (FP)
var finalPermutation = []int{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25,
}

// Tabela de expansão E
var expansionTable = []int{
	32, 1, 2, 3, 4, 5,
	4, 5, 6, 7, 8, 9,
	8, 9, 10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 31, 32, 1,
}

// Tabela de permutação P
var pPermutation = []int{
	16, 7, 20, 21, 29, 12, 28, 17,
	1, 15, 23, 26, 5, 18, 31, 10,
	2, 8, 24, 14, 32, 27, 3, 9,
	19, 13, 30, 6, 22, 11, 4, 25,
}

// S-boxes (8 caixas de substituição de 4x16)
var sBoxes = [8][4][16]int{
	{ // S1
		{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
		{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
		{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
		{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
	},
	{ // S2
		{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
		{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
		{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
		{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
	},
	{ // S3
		{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
		{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
		{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
		{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
	},
	{ // S4
		{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
		{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
		{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
		{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
	},
	{ // S5
		{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
		{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
		{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
		{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
	},
	{ // S6
		{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
		{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
		{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
		{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
	},
	{ // S7
		{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
		{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
		{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
		{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
	},
	{ // S8
		{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
		{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
		{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
		{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
	},
}

// Tabelas para geração das chaves
var pc1 = []int{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4,
}

var pc2 = []int{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53,
	46, 42, 50, 36, 29, 32,
}

var leftShifts = []int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

// Estrutura principal do DES
type DES struct {
	keys [16]uint64
}

// NewDES cria uma nova instância do DES com a chave fornecida
func NewDES(key []byte) (*DES, error) {
	if len(key) != 8 {
		return nil, fmt.Errorf("chave deve ter exatamente 8 bytes")
	}
	
	des := &DES{}
	des.generateKeys(key)
	return des, nil
}

// Converte slice de bytes para uint64
func bytesToUint64(b []byte) uint64 {
	var result uint64
	for i := 0; i < 8; i++ {
		result = (result << 8) | uint64(b[i])
	}
	return result
}

// Converte uint64 para slice de bytes
func uint64ToBytes(val uint64) []byte {
	result := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		result[i] = byte(val & 0xFF)
		val >>= 8
	}
	return result
}

// Aplica uma permutação
func permute(input uint64, table []int, inputLen int) uint64 {
	var output uint64
	for i, pos := range table {
		bit := (input >> (inputLen - pos)) & 1
		output |= bit << (len(table) - 1 - i)
	}
	return output
}

// Rotação circular à esquerda
func leftRotate(val uint64, bits, size int) uint64 {
	mask := (uint64(1) << size) - 1
	val &= mask
	return ((val << bits) | (val >> (size - bits))) & mask
}

// Gera as 16 chaves de round
func (d *DES) generateKeys(key []byte) {
	keyUint64 := bytesToUint64(key)
	
	// PC-1 permutation
	key56 := permute(keyUint64, pc1, 64)
	
	// Divide em duas metades
	left := (key56 >> 28) & 0xFFFFFFF
	right := key56 & 0xFFFFFFF
	
	// Gera as 16 chaves
	for i := 0; i < 16; i++ {
		// Rotação
		left = leftRotate(left, leftShifts[i], 28)
		right = leftRotate(right, leftShifts[i], 28)
		
		// Combina e aplica PC-2
		combined := (left << 28) | right
		d.keys[i] = permute(combined, pc2, 56)
	}
}

// Função f do algoritmo DES
func feistelFunction(right uint32, key uint64) uint32 {
	// Expansão E
	expanded := uint64(permute(uint64(right), expansionTable, 32))
	
	// XOR com a chave
	xored := expanded ^ key
	
	// S-boxes
	var sBoxOutput uint32
	for i := 0; i < 8; i++ {
		// Extrai 6 bits
		sixBits := (xored >> (42 - 6*i)) & 0x3F
		
		// Calcula linha e coluna para S-box
		row := ((sixBits & 0x20) >> 4) | (sixBits & 0x01)
		col := (sixBits & 0x1E) >> 1
		
		// Aplica S-box
		sBoxValue := sBoxes[i][row][col]
		sBoxOutput |= uint32(sBoxValue) << (28 - 4*i)
	}
	
	// Permutação P
	return uint32(permute(uint64(sBoxOutput), pPermutation, 32))
}

// Encripta um bloco de 8 bytes
func (d *DES) encryptBlock(block []byte) []byte {
	if len(block) != 8 {
		panic("Bloco deve ter 8 bytes")
	}
	
	data := bytesToUint64(block)
	
	// Permutação inicial
	data = permute(data, initialPermutation, 64)
	
	// Divide em duas metades
	left := uint32(data >> 32)
	right := uint32(data & 0xFFFFFFFF)
	
	// 16 rounds
	for i := 0; i < 16; i++ {
		newRight := left ^ feistelFunction(right, d.keys[i])
		left = right
		right = newRight
	}
	
	// Combina (note a troca de posições)
	combined := (uint64(right) << 32) | uint64(left)
	
	// Permutação final
	result := permute(combined, finalPermutation, 64)
	
	return uint64ToBytes(result)
}

// Decripta um bloco de 8 bytes
func (d *DES) decryptBlock(block []byte) []byte {
	if len(block) != 8 {
		panic("Bloco deve ter 8 bytes")
	}
	
	data := bytesToUint64(block)
	
	// Permutação inicial
	data = permute(data, initialPermutation, 64)
	
	// Divide em duas metades
	left := uint32(data >> 32)
	right := uint32(data & 0xFFFFFFFF)
	
	// 16 rounds (chaves em ordem reversa)
	for i := 15; i >= 0; i-- {
		newRight := left ^ feistelFunction(right, d.keys[i])
		left = right
		right = newRight
	}
	
	// Combina (note a troca de posições)
	combined := (uint64(right) << 32) | uint64(left)
	
	// Permutação final
	result := permute(combined, finalPermutation, 64)
	
	return uint64ToBytes(result)
}

// Aplica padding PKCS7
func addPadding(data []byte) []byte {
	padding := 8 - (len(data) % 8)
	padText := make([]byte, len(data)+padding)
	copy(padText, data)
	for i := len(data); i < len(padText); i++ {
		padText[i] = byte(padding)
	}
	return padText
}

// Remove padding PKCS7
func removePadding(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	padding := int(data[len(data)-1])
	if padding < 1 || padding > 8 || padding > len(data) {
		return data
	}
	return data[:len(data)-padding]
}

// Encrypt encripta uma mensagem completa
func (d *DES) Encrypt(plaintext []byte) []byte {
	paddedData := addPadding(plaintext)
	ciphertext := make([]byte, 0, len(paddedData))
	
	for i := 0; i < len(paddedData); i += 8 {
		block := paddedData[i : i+8]
		encryptedBlock := d.encryptBlock(block)
		ciphertext = append(ciphertext, encryptedBlock...)
	}
	
	return ciphertext
}

// Decrypt decripta uma mensagem completa
func (d *DES) Decrypt(ciphertext []byte) []byte {
	if len(ciphertext)%8 != 0 {
		panic("Texto cifrado deve ser múltiplo de 8 bytes")
	}
	
	plaintext := make([]byte, 0, len(ciphertext))
	
	for i := 0; i < len(ciphertext); i += 8 {
		block := ciphertext[i : i+8]
		decryptedBlock := d.decryptBlock(block)
		plaintext = append(plaintext, decryptedBlock...)
	}
	
	return removePadding(plaintext)
}

func main() {
	fmt.Println("=== Implementação DES (Data Encryption Standard) ===\n")
	
	// Exemplo de uso
	key := []byte("CHAVE123") // 8 bytes
	plaintext := "Esta é uma mensagem secreta para demonstrar o algoritmo DES!"
	
	fmt.Printf("Chave: %s\n", string(key))
	fmt.Printf("Texto original: %s\n", plaintext)
	
	// Cria instância do DES
	des, err := NewDES(key)
	if err != nil {
		log.Fatal(err)
	}
	
	// Encripta
	ciphertext := des.Encrypt([]byte(plaintext))
	fmt.Printf("Texto cifrado (hex): %s\n", hex.EncodeToString(ciphertext))
	
	// Decripta
	decrypted := des.Decrypt(ciphertext)
	fmt.Printf("Texto decriptado: %s\n", string(decrypted))
	
	// Verifica se a decriptação foi bem-sucedida
	if string(decrypted) == plaintext {
		fmt.Println("\n✅ Sucesso! A mensagem foi cifrada e decifrada corretamente.")
	} else {
		fmt.Println("\n❌ Erro! A decriptação falhou.")
	}
	
	// Demonstração interativa
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Demonstração interativa:")
	fmt.Print("Digite uma mensagem para cifrar: ")
	
	var userMessage string
	fmt.Scanln(&userMessage)
	
	if userMessage != "" {
		userCiphertext := des.Encrypt([]byte(userMessage))
		userDecrypted := des.Decrypt(userCiphertext)
		
		fmt.Printf("\nMensagem original: %s\n", userMessage)
		fmt.Printf("Cifrada (hex): %s\n", hex.EncodeToString(userCiphertext))
		fmt.Printf("Decifrada: %s\n", string(userDecrypted))
	}
}
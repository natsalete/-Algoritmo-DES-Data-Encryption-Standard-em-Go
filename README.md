# DES (Data Encryption Standard) Implementation in Go

## Project Description

This project implements the **DES (Data Encryption Standard)** algorithm from scratch in Go, without using pre-built cryptography libraries. DES is a symmetric encryption algorithm that uses a block cipher, operating on 64-bit blocks with a 56-bit effective key (64-bit with parity).

## How the DES Algorithm Works

DES is based on the Feistel Lattice and performs the following main steps:

### 1. Key Generation
- The 64-bit key is reduced to 56 bits through PC-1 permutation
- It is divided into two 28-bit halves
- It generates 16 48-bit subkeys through rotations and PC-2 permutation

### 2. Encryption Process
- Initial Permutation (IP): Rearranges the 64 bits of the data block
- 16 Feistel Rounds: Each round applies:
- Division of the block into two halves (L and R)
- Function f applied to the right half with the round's subkey
- XOR of the result with the left half
- Swap the halves
- Final Permutation (FP): Final rearrangement of the bits

### 3. Function f (Kernel) (DES)
- **E Expansion**: Expands 32 bits to 48 bits
- **XOR**: With the 48-bit subkey
- **S-boxes**: 8 substitution boxes that reduce 48 bits to 32 bits
- **P Permutation**: Rearranges the resulting 32 bits

### 4. Decryption
- Same process as encryption, but with the subkeys applied in reverse order

## Implementation Features

### Implemented Features
- ✅ Complete generation of all 16 subkeys
- ✅ All permutations (IP, FP, E, P)
- ✅ All 8 S-boxes of the DES standard
- ✅ Full Feistel functionality
- ✅ PKCS#7 padding for messages of any length
- ✅ ECB (Electronic Codebook) mode
- ✅ Simple interface for Encryption/Decryption

### Data Structures
- **Permutation Tables**: IP, FP, E, P, PC-1, PC-2
- **S-boxes**: 8 4×16 substitution tables
- **Rotations**: Number of rotations per round

## How to Run

### Prerequisites
- Go 1.19 or higher installed
- Operating System: Windows, Linux, or macOS

### Installation and Running

1. **Clone the repository:**
```bash
git clone <REPOSITORY_URL>
cd des-cipher-go
```

2. **Run the program:**
```bash
go run main.go
```

3. **Or compile and run:**
```bash
go build -o des-cipher main.go
./des-cipher
```

## Example Usage

```go
package main

import (
"fmt"
"encoding/hex"
)

func main() {
// 8-byte (64-bit) key
key := []byte("KEY123")

// Message to be encrypted
plaintext := "Secret message!"

// Create an instance of DES
des, err := NewDES(key)
if err != nil {
panic(err)
}

// Encrypt the message
ciphertext := des.Encrypt([]byte(plaintext))
fmt.Printf("Encrypted: %s\n", hex.EncodeToString(ciphertext))

// Decrypt the message
decrypted := des.Decrypt(ciphertext)
fmt.Printf("Decrypted: %s\n", string(decrypted))
}
```

## Demonstration

The program includes:
- **Automatic example**: Demonstrates the encryption of a predefined message
- **Interactive mode**: Allows the user to enter a custom message
- **Verification**: Confirms that the decryption returns the plaintext
- **Hexadecimal display**: Displays the ciphertext in hexadecimal format

## Security and Limitations

### ⚠️ Important Notice
This implementation is **for educational purposes only**. DES is considered **insecure** for production use due to:
- Only 56-bit key (easily cracked by brute force)
- Known vulnerabilities
- Replaced by more secure algorithms (AES, 3DES)

### Implementation Limitations
- **ECB mode only**: Does not implement more secure modes (CBC, CTR)
- **No authentication**: Does not include integrity checking
- **Performance**: Optimized for clarity, not speed

## Project Structure

```
des-cipher-go/
├── main.go # Main code with complete implementation
├── README.md # This documentation
```

## Demonstrated Concepts

- **Symmetric Encryption**: Same key for encryption and decryption
- **Block Cipher**: Processing in fixed 64-bit blocks
- **Network Feistel's**: Fundamental Structure of DES
- **Permutations**: Controlled Bit Rearrangement
- **S-Boxes**: Nonlinear Substitution for Security
- **Key Generation**: Derivation of Multiple Subkeys
- **Padding**: Handling Variable-Length Messages

## References

- FIPS 46-3: Data Encryption Standard (DES)
- "Applied Cryptography" - Bruce Schneier
- NIST Special Publication 800-67: Recommendation for the Triple Data Encryption Algorithm (TDEA) Block Cipher

## Author

Developed with ❤️ to demonstrate the fundamentals of symmetric encryption and block cipher algorithms.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

<details>
<summary>🇧🇷 Versão em Português</summary>

⭐ **If this project was useful to you, don't forget to give it a star.

# Implementação DES (Data Encryption Standard) em Go

## Descrição do Projeto

Este projeto implementa o algoritmo **DES (Data Encryption Standard)** do zero em Go, sem o uso de bibliotecas de criptografia prontas. O DES é um algoritmo de criptografia simétrica que utiliza cifra por blocos, operando em blocos de 64 bits com uma chave de 56 bits efetivos (64 bits com paridade).

## Como Funciona o Algoritmo DES

O DES é baseado na **Rede de Feistel** e realiza os seguintes passos principais:

### 1. Geração das Chaves
- A chave de 64 bits é reduzida para 56 bits através da permutação PC-1
- É dividida em duas metades de 28 bits cada
- Gera 16 subchaves de 48 bits através de rotações e permutação PC-2

### 2. Processo de Cifragem
- **Permutação Inicial (IP)**: Reorganiza os 64 bits do bloco de dados
- **16 Rounds de Feistel**: Cada round aplica:
  - Divisão do bloco em duas metades (L e R)
  - Função f aplicada à metade direita com a subchave do round
  - XOR do resultado com a metade esquerda
  - Troca das metades
- **Permutação Final (FP)**: Reorganização final dos bits

### 3. Função f (Núcleo do DES)
- **Expansão E**: Expande 32 bits para 48 bits
- **XOR**: Com a subchave de 48 bits
- **S-boxes**: 8 caixas de substituição que reduzem de 48 para 32 bits
- **Permutação P**: Reorganiza os 32 bits resultantes

### 4. Decifragem
- Mesmo processo da cifragem, mas com as subchaves aplicadas em ordem reversa

## Características da Implementação

### Recursos Implementados
- ✅ Geração completa das 16 subchaves
- ✅ Todas as permutações (IP, FP, E, P)
- ✅ Todas as 8 S-boxes do padrão DES
- ✅ Função de Feistel completa
- ✅ Padding PKCS#7 para mensagens de qualquer tamanho
- ✅ Modo ECB (Electronic Codebook)
- ✅ Interface simples para cifragem/decifragem

### Estruturas de Dados
- **Tabelas de Permutação**: IP, FP, E, P, PC-1, PC-2
- **S-boxes**: 8 tabelas de substituição 4×16
- **Rotações**: Quantidade de rotações por round

## Como Executar

### Pré-requisitos
- Go 1.19 ou superior instalado
- Sistema operacional: Windows, Linux ou macOS

### Instalação e Execução

1. **Clone o repositório:**
```bash
git clone <URL_DO_REPOSITORIO>
cd des-cipher-go
```

2. **Execute o programa:**
```bash
go run main.go
```

3. **Ou compile e execute:**
```bash
go build -o des-cipher main.go
./des-cipher
```

## Exemplo de Uso

```go
package main

import (
    "fmt"
    "encoding/hex"
)

func main() {
    // Chave de 8 bytes (64 bits)
    key := []byte("CHAVE123")
    
    // Mensagem a ser cifrada
    plaintext := "Mensagem secreta!"
    
    // Cria instância do DES
    des, err := NewDES(key)
    if err != nil {
        panic(err)
    }
    
    // Cifra a mensagem
    ciphertext := des.Encrypt([]byte(plaintext))
    fmt.Printf("Cifrado: %s\n", hex.EncodeToString(ciphertext))
    
    // Decifra a mensagem
    decrypted := des.Decrypt(ciphertext)
    fmt.Printf("Decifrado: %s\n", string(decrypted))
}
```

## Demonstração

O programa inclui:
- **Exemplo automático**: Demonstra a cifragem de uma mensagem pré-definida
- **Modo interativo**: Permite ao usuário inserir uma mensagem personalizada
- **Verificação**: Confirma se a decifragem retorna o texto original
- **Exibição hexadecimal**: Mostra o texto cifrado em formato hexadecimal

## Segurança e Limitações

### ⚠️ Aviso Importante
Esta implementação é **exclusivamente educacional**. O DES é considerado **inseguro** para uso em produção devido a:
- Chave de apenas 56 bits (facilmente quebrada por força bruta)
- Vulnerabilidades conhecidas
- Substituído por algoritmos mais seguros (AES, 3DES)

### Limitações da Implementação
- **Modo ECB apenas**: Não implementa modos mais seguros (CBC, CTR)
- **Sem autenticação**: Não inclui verificação de integridade
- **Performance**: Otimizada para clareza, não para velocidade

## Estrutura do Projeto

```
des-cipher-go/
├── main.go          # Código principal com implementação completa
├── README.md        # Esta documentação
```

## Conceitos Demonstrados

- **Criptografia Simétrica**: Mesma chave para cifrar e decifrar
- **Cifra por Blocos**: Processamento em blocos fixos de 64 bits
- **Rede de Feistel**: Estrutura fundamental do DES
- **Permutações**: Reorganização controlada de bits
- **S-boxes**: Substituição não-linear para segurança
- **Geração de Chaves**: Derivação de múltiplas subchaves
- **Padding**: Tratamento de mensagens de tamanho variável

## Referências

- FIPS 46-3: Data Encryption Standard (DES)
- "Applied Cryptography" - Bruce Schneier
- NIST Special Publication 800-67: Recommendation for the Triple Data Encryption Algorithm (TDEA) Block Cipher

## Autor

Desenvolvido com ❤️ para demonstrar os fundamentos da criptografia simétrica e algoritmos de cifra por blocos.

## Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

---

⭐ **Se este projeto foi útil para você, não esqueça de dar uma estrela!** ⭐

</details>


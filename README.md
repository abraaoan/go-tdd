# 📝 Todo List CLI (Go)

Aplicação de gerenciamento de tarefas com autenticação de usuário, escrita em Go, com arquitetura limpa, TDD e persistência em PostgreSQL. Interface de uso via terminal (linha de comando).

---

## 📦 Funcionalidades

- [x] Criar usuário
- [x] Login com geração de token
- [x] Criar tarefa vinculada ao usuário
- [x] Listar tarefas do usuário logado
- [x] Completar tarefa
- [x] Deletar tarefa
- [x] Interface interativa via terminal (CLI)

---

## 🧱 Arquitetura

O projeto segue o padrão **Ports and Adapters (Hexagonal Architecture)** com separação clara entre:

- `entity/` — entidades de domínio puras
- `usecase/` — regras de negócio (application services)
- `repository/` — interfaces para persistência
- `provider/` — abstrações como geração de tokens
- `adapter/postgres/` — implementações de banco
- `cmd/` — interface de linha de comando (CLI)
- `internal/testes/` — testes organizados por comportamento

---

## 🚀 Como rodar

### 1. Clone o projeto

```bash
git clone https://github.com/seuusuario/todo-list.git
cd todo-list
```

### 2. Configure o banco de dados PostgreSQL

Crie um banco e a tabela `User` (ou `"User"` com aspas se preferir):

```sql
CREATE TABLE "User" (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL
);

CREATE TABLE task (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  status BOOLEAN NOT NULL DEFAULT false,
  user_id INTEGER REFERENCES "User"(id)
);
```

### 3. Configure `.env` ou conexão manual

Você pode configurar manualmente no `main.go` ou usar uma lib como `godotenv`.

### 4. Rode o app

```bash
go run main.go
```

Você verá algo assim:

```
O que você quer fazer?
1. Criar usuário
2. Login
3. Criar tarefa
4. Listar tarefas
0. Sair
> 
```

---

## 🧪 Testes

O projeto usa:

- `testing`
- `github.com/stretchr/testify`
- `github.com/golang/mock/gomock`

### Rodar todos os testes:

```bash
go test ./...
```

Os testes seguem o ciclo TDD e estão organizados por caso de uso.

---

## 🔐 Sobre autenticação

- Ao criar usuário ou fazer login, é gerado um "token fake" com base no `userID`.
- O token ainda **não é JWT**, mas o sistema já está pronto para plugar um `JwtTokenProvider` no futuro.
- As tarefas são sempre associadas a um usuário via `userID`.

---

## 📚 Exemplo de uso

```bash
> Criar usuário
Email: joao@email.com
Senha: 123456
Token: token-1-1725407891404

> Login
Email: joao@email.com
Senha: 123456
Token: token-1-1725407891441

> Criar tarefa
Título: Estudar TDD
Tarefa criada com ID 1

> Listar tarefas
ID:     1
Título: Estudar TDD
Status: ❌ Pendente
Dono:   1
------------------------------
```

---

## ✨ Futuras melhorias

- [ ] Adicionar hash de senha com bcrypt
- [ ] Implementar autenticação JWT real
- [ ] Interface HTTP com os mesmos usecases
- [ ] Filtrar tarefas concluídas
- [ ] Multi-usuário com sessão persistente

---

## 🧑‍💻 Autor

Desenvolvido por abraaoan com arquitetura limpa, testes e Go idiomático.

---

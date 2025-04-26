2025-04-24 15:42

Status: #project

Tags: #git #github #rest_and_go_pet #terminal #code #bash

# Git
- Git init is a just easy weight screens 
- Git branch also easy weight. Its just a link on definite commit. U can create many branches and them do nothing to memory.
- HEAD - symbolical name of the [[Book translator#^8f8559|current]] commit. HEAD always indicates on ur last commit from ur local tree. ^6c5aea


# Steps
## Настройка SSH и подключение к репу
1. Сгенерировать SSH-ключ
```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```
2. Добавить ключ в GitHub
Скопировать содержимое `~/.ssh/id_ed25519.pub`.
В GitHub: `Settings` → `SSH and GPG keys` → `New SSH key`.
3. Клонировать реп и сменить директорию
```bash
git clone git@github.com:Danyarbrg/securelogAPI.git
cd securelogAPI
```
4. Конфиг
```bash
git config --global user.name "<ваше_имя>"
git config --global user.email "<адрес_почты@email.com>"
```
## Работа с Git
- Check status
```bash
git status
```
- Create repository %% `cd` on your directory%%
```bash
git init
```
- U should create git branch and switch on it 
```bash
git branch <branch_name> # create
git checkout <branch_name> # change branch
git checkout -b <branch_name> # create and switch by one command 
```
- See all branches
```bash
git branch
```
- Add all files or add smth definite
```bash
git add .
# or definite
git add <file_name>
```
- Delete changes in rep
```bash
git restore filename
```
- After create commit. necessarily dont forget about commenting 
```bash
git commit filename -m "comment"
```
- To connect branches u have to merge ***bugFix*** and ***main*** (ur main branch is main)
  after u need to switch branch on ***bugFix*** and merge again
```bash
git merge bugFix
git checkout bugFix 
git merge main
```
- Also u can connect branches with `merge` for a cleaner connections
```bash
git merge main # bugFix to main
```
- If u want to jump between chapters in branch and change smth im past commit, u have to 
- To switch on other hash-commit u can write only first pair of letters
```bash
git check fed2 # where fed2 is fed2da64c0efc5293610bdd892f82a58e8cbc5d8 == hash name of commit
```
- But u can move between commits with easier commands 
	- Move back `^`
	- Move back on several meanings `~<num>`
```bash
git checkout HEAD~4
```
- When we use `git checkout main^` we move to main
commit (parent) c0 -> c1 -> c2(пргынем сюда) -> c3 (HEAD)
`git checkout main^^` we move to [[Book translator^123|progenitor]] (прародитель)
- We can also use [[Git#^6c5aea|HEAD]] to move in commits 
```bash
git checkout fed2 # commit 2
git checkout HEAD^ # commit 1
git checkout HEAD^ # commit 0
```
---
- To see changes and differences in files 
```bash
git diff
```
- Before starting work u need to write this code to download all commites and changes 
```bash
git pull --rebase
```
- See info about hashes and names of branches, commites info. Also u can check where is [[Git#^6c5aea|HEAD]] now
```bash
git log
```
- See graph of commits 
```bash
git log --oneline --graph --decorate --all
```
# Terms
# Footnote
# References
[Интерактивный сайт с изучением git](https://learngitbranching.js.org/?locale=ru_RU&demo=)
[Хабр статья](https://habr.com/ru/articles/767424/)
[lazygit](https://habr.com/ru/companies/flant/articles/712874/)

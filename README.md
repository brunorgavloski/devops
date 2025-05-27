# 3 git
- git init
- git add 'file'
- git commit -m
- git status
- git log
- git log --author='name'
- git clone -b main https://url-of-repo
- git pull - sync repo remote with local 
- git push - sync local with remote
- git push -u origin main
# 4 GitHub
Git(local) ->Push-> GitHub(remote)
Git(local) <-Pull<- GitHub(remote)
Add .git ignore
- Local configure
- Go to account settings - ssh and gpts keys - new SSH key - Give a name - place pub key - test the conection
- ssh-keygen -t ed25519 -C "seu_email@example.com"
- eval "$(ssh-agent -s)"
- ssh-add ~/.ssh/id_ed25519
- cat ~/.ssh/id_ed25519.pub
- ssh -T git@github.com

'''
# TMUX
Ctrl + b, "  → para dividir na horizontal  
Ctrl + b, %  → para dividir na vertical
Ctrl + b (prefixo padrão), depois:

← → vai para o painel à esquerda.

→ → vai para o painel à direita.

↑ → vai para o painel acima.

↓ → vai para o painel abaixo.


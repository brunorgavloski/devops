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
- git branch - list the branch
- git branch 'name' - create a new branch
- git switch 'name'
- git merge 'develop'
- git pull --no-rebase
- git reset - rollback - remove the last commit
- git revert - create a new commit at end of the chain to 'cancel' the change
- git log --oneline
- git reset --soft 'id_commit' - no change de files 'soft'
- git reset --hard 'id_commit'
- git diff 'file_name' - show the changes of the file
- git diff - compara stage area
- git diff - 'id_01' 'id_02'
- git branch -a
- git branch -m 'go' 'go1' - rename the branch
- git branch -d 'name' - delet the branch
- git stash


'''
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
# Jenkins
- docker run -p 8080:8080  jenkins/jenkins:lts - Não persistente
- docker run -d -p 8080:8080  jenkins/jenkins:lts - Background
- docker logs 'id_container' - Pegar a senha
- docker run -d -p 8080:8080 -v /jenkins_data:/var/jenkins_home jenkins/jenkins:lts











'''
# TMUX
Ctrl + b, "  → para dividir na horizontal  
Ctrl + b, %  → para dividir na vertical
Ctrl + b (prefixo padrão), depois:
← → vai para o painel à esquerda.
→ → vai para o painel à direita.
↑ → vai para o painel acima.
↓ → vai para o painel abaixo.


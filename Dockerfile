FROM fedora:42

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin v3.44.0

RUN dnf --setopt=install_weak_deps=False install jq ripgrep git -y

WORKDIR /app

CMD ["task", "--version"]

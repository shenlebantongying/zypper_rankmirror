Sample output

```bash
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed -v
[sudo] password for root: 
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed removerepo -a
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed ar -c  http://download.opensuse.org/tumbleweed/repo/oss/ Tumbleweed
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed refresh
Refresh test: openSUSE offical duration:19.175762726s
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed install --no-recommends -y --download-only filesystem
Install test: openSUSE offical duration:1.739925274s
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed removerepo -a
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed ar -c http://mirrors.tuna.tsinghua.edu.cn/opensuse/tumbleweed/repo/oss/ Tumbleweed
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed refresh
Refresh test: 清华 duration:17.399546997s
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed install --no-recommends -y --download-only filesystem
Install test: 清华 duration:2.512386229s
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed removerepo -a
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed ar -c http://mirror.bjtu.edu.cn/opensuse/tumbleweed/repo/oss/ Tumbleweed
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed refresh
Refresh test: 北交大 duration:21.787947682s
[->] sudo zypper --no-gpg-checks --root=/home/slbtty/.cache/tumbleweed install --no-recommends -y --download-only filesystem
Install test: 北交大 duration:2.749009382s

```
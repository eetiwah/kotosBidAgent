For whatever reason, you must run this statement so that the environment will be able to find tor and associated libraries, I will add to the codebase later to automate this process

export LD_LIBRARY_PATH=/home/rich/Projects/kotos/kotosBidAgent/agent/tor/linux/dependencies:$LD_LIBRARY_PATH
/home/rich/Projects/kotos/kotosBidAgent/agent/tor/linux/dependencies/tor --version
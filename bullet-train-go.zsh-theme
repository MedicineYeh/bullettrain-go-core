# Prompt previous command execution time
preexec() {
  cmd_timestamp=`date +%s`
}

precmd() {
  [[ $BULLETTRAIN_CAR_EXEC_TIME_SHOW == false ]] && return

  local stop=`date +%s`
  local start=${cmd_timestamp:-$stop}
  let BULLETTRAIN_last_exec_duration=$stop-$start
  cmd_timestamp=''
}

BULLETTRAIN_script_dir=$(dirname $0)
if [[ $OSTYPE == *darwin* ]]; then
    # MAC OS
    PROMPT='$(${BULLETTRAIN_script_dir}/releases/bullettrain.darwin-amd64 $? $BULLETTRAIN_last_exec_duration)'
else
    # Linux
    case $(uname -m) in
    x86_64)
        PROMPT='$(${BULLETTRAIN_script_dir}/releases/bullettrain.linux-amd64 $? $BULLETTRAIN_last_exec_duration)'
    ;;
    armv6*)
        PROMPT='$(${BULLETTRAIN_script_dir}/releases/bullettrain.linux-armv6 $? $BULLETTRAIN_last_exec_duration)'
    ;;
    armv7*)
        PROMPT='$(${BULLETTRAIN_script_dir}/releases/bullettrain.linux-armv7 $? $BULLETTRAIN_last_exec_duration)'
    ;;
    *)
        PROMPT='$(${BULLETTRAIN_script_dir}/releases/bullettrain.linux-arm64 $? $BULLETTRAIN_last_exec_duration)'
    ;;
    esac
fi

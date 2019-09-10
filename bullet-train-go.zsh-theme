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

BULLETTRAIN_script_dir=$(dirname ${ZSH_CUSTOM}/themes/${ZSH_THEME})
PROMPT='$(${BULLETTRAIN_script_dir}/releases/bullettrain.linux-amd64 $? $BULLETTRAIN_last_exec_duration)'

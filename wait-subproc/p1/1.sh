
 pross=()
 ./2.sh &
 pross=(${pross[@]} $!)
 ./2.sh &
 pross=(${pross[@]} $!)
 ./2.sh &
 pross=(${pross[@]} $!)
 ./2.sh &
 pross=(${pross[@]} $!)
 for p in ${pross[@]} ; do 
 	wait ${p}
	echo ${p} done status: $?
 done

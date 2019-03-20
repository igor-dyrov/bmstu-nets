const array = [5, 7, 1, 0, 9, 11, 3];

const mergeSort = (arr) => {
   if(arr.length < 2)
      return arr;
   const mid = Math.floor(arr.length/2),
       left = arr.slice(0,mid),
       right = arr.slice(mid);
   return merge(mergeSort(left),mergeSort(right));
}

const merge = (left, right) => {
	const result = [];
    let  lLen = left.length,
      rLen = right.length,
      l = 0,
      r = 0;
  while(l < lLen && r < rLen){
     if(left[l] < right[r]){
       result.push(left[l++]);
     }
     else{
       result.push(right[r++]);
    }
  }  
  return result.concat(left.slice(l)).concat(right.slice(r));
}


console.log(mergeSort(array));

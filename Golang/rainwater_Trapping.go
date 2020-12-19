

//Function defining Rain Trap with arguments consisting of height
func rain_trap(height []int) int {
	var res int                      //To store max water that can be stored
	for i := 1; i < len(height)-1; i++ {                  //Using loop for traversing throughout the array or heights
		var leftMax, rightMax = height[i], height[i]
		for l, r := i-1, i+1; l >= 0 || r < len(height); l, r = l-1, r+1 {
			//Using condition to check if l>=0
			if l >= 0 {
				leftMax = max(leftMax, height[l])
			}
				//Using condition to check if r< length of the height
			if r < len(height) {
				rightMax = max(rightMax, height[r])
			}
		}
		res += min(leftMax, rightMax) - height[i]
	}
	return res



	//SAMPLE I/0
	
	//Input
    //arr = [0, 1, 0, 2, 1, 0, 
           //1, 3, 2, 1, 2, 1]; 
    //n = len(arr); 
     
	//Println(maxWater(arr, n)); 
	//output
	//6
     
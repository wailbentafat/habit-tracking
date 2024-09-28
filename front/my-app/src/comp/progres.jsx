const Progress = () => {
   //fetch level
   //fetch xp out of 100
   let max=100
   let Level=7
   let Xp=75 
   let remainingXp=max-Xp
   let percentage= (Xp/max)*100
    return (
        <div className="w-64 h-screen bg-white p-2 shadow-md">
            <h1 className="text-2xl font-bold *:hover text-center" >Level {Level}</h1>

            <div className="mt-3">
                <div className="text-sm  text-center tracking-wide text-gray-500 uppercase font-medium ">{Xp} XP / {max} XP</div>
                <div className=" mt-2 w-full bg-gray-200 rounded-full h-2">
            <div
                className="bg-black h-2 rounded-full"
                style={{ width: `${percentage}%` }}
            />
        </div>
        <div className="text-sm mt-2  text-center tracking-wide text-gray-500 uppercase font-medium ">{remainingXp} XP remaining</div>
    
             </div>
            </div>
                
              
                
               
                
                

    )
}
export default Progress
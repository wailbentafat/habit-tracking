import Progress from "./progres"

export const SideBar = () => {

    return (
           <div className="w-64 h-screen bg-white p-6 shadow-md">
            <div className="h-32 overflow-hidden">
            <Progress />
            </div>
           
            <div className="mt-6">
                <h3 className="text-l font-bold">Quick States</h3>
                <ul className="mt-2 text-sm text-gray-500">
                    <li className="py-2 pl-6 ">Habits Completed: 3/5 </li>
                    <li className="py-2 pl-6">Weekly Streak: 5 days</li>
                    <li className="py-2 pl-6">Monthly Progress: 85%</li>

                </ul>
              
            </div>
            
            
        </div>
    )
}
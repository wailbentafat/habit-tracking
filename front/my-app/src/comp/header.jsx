import { FaBell, FaUser, FaCog } from 'react-icons/fa';


export const Header = () => {
    return (
        <div  className=" h-20 flex bg-white p-2 shadow-md w-full">
            <h1 className="text-2xl mt-3  ml-2 text-left font-extrabold ">HabitTracker</h1>
            <div className="ml-auto mr-5 flex items-center  space-x-10"> 
                
                    <FaBell className=" ml-6 text-gray-500 hover:text-gray-700 focus:text-gray-700 focus:outline-none size-5"/>
                    <FaUser className="ml-6 text-gray-500 hover:text-gray-700 focus:text-gray-700 focus:outline-none  size-5" />
                    <FaCog className="ml-6 text-gray-500 hover:text-gray-700 focus:text-gray-700 focus:outline-none  size-5" />
                

            </div>
        </div>
        )}
        
        
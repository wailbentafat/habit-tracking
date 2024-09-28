import React from 'react';
import { Header } from './header';
import { SideBar } from './sidbr';
import { Habits } from './habits';
//categorie 

export const Layout = () => {
    return (
        <div style={{ display: 'flex', height: '100vh' }}>
            <div style={{ width: '250px' }}> 
          
                <SideBar />
            </div>
            <div className="flex-1 flex flex-col overflow-y-auto bg-gray-100">
                <Header className="w-full" /> 

                <div className="p-6 flex ">
                    
                    <h2 className='text-xl font-semibold  ml-2 mt-2 '>My Habits </h2>
                        <button className="flex ml-auto mt-1 bg-black  text-white font-normal py-2 px-4 rounded-md  hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"> +  Add New Habit</button>
                   
                        
                    
                </div>
                <div className="p-3 flex-col ml-4 ">
                <button className='bg-white px-6 py-2 border border-gray-300 rounded-md font-normal '> Fitness  </button>
                <button className='ml-4 bg-white px-6 py-2 border border-gray-300 rounded-md font-normal'> Health</button>
                <button className='ml-4 bg-white px-6 py-2 border border-gray-300 rounded-md font-normal'>  Learning</button>
                <button className='ml-4 bg-white px-6 py-2 border border-gray-300 rounded-md font-normal'> Productivity</button>
                <button className='ml-4 bg-white px-6 py-2 border border-gray-300 rounded-md font-normal'> Wellness</button>
                </div>
                <div className=' space-y-4 '>
                    <Habits />
                </div>
            </div>
        </div>
    );
};

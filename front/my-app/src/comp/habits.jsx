import React from 'react';
export const Habits = () => {
    
    const habitsData = {
        message: "habits retrieved",
        habits: [
            {
                id: 1,
                userid: 123,
                categoryId: 1,
                name: "Exercise",
                createdat: "2024-09-28T12:00:00Z",
                updatedAt: "2024-09-28T12:00:00Z",
                reminders: [
                    {
                        id: 1,
                        habitid: 1,
                        date: "2024-09-28T09:00:00Z",
                        status: "active"
                    }
                ],
                goals: [
                    {
                        id: 1,
                        habitid: 1,
                        target: 5,
                        current: 3,
                        streak: 1,
                        date: "2024-09-28T12:00:00Z",
                        updatedAt: "2024-09-28T12:00:00Z"
                    }
                ]
            },
            {
                id: 2,
                userid: 124,
                categoryId: 2,
                name: "Reading",
                createdat: "2024-09-28T12:00:00Z",
                updatedAt: "2024-09-28T12:00:00Z",
                reminders: [
                    {
                        id: 2,
                        habitid: 2,
                        date: "2024-09-28T10:00:00Z",
                        status: "active"
                    }
                ],
                goals: [
                    {
                        id: 2,
                        habitid: 2,
                        target: 30,
                        current: 15,
                        streak: 2,
                        date: "2024-09-28T12:00:00Z",
                        updatedAt: "2024-09-28T12:00:00Z"
                    }
                ]
            },
            {
                id: 3,
                userid: 125,
                categoryId: 3,
                name: "Meditation",
                createdat: "2024-09-28T12:00:00Z",
                updatedAt: "2024-09-28T12:00:00Z",
                reminders: [
                    {
                        id: 3,
                        habitid: 3,
                        date: "2024-09-28T08:00:00Z",
                        status: "active"
                    }
                ],
                goals: [
                    {
                        id: 3,
                        habitid: 3,
                        target: 10,
                        current: 7,
                        streak: 1,
                        date: "2024-09-28T12:00:00Z",
                        updatedAt: "2024-09-28T12:00:00Z"
                    }
                ]
            }
        ]
    };

    const percentage = (habitsData.habits[0].goals[0].current / habitsData.habits[0].goals[0].target) * 100;
  

    return (
        <div className="flex flex-col space-y-4 ml-7 mr-8 mt-3">
            {habitsData.habits.map((habit) => (
                <div key={habit.id} className=" flex-1 min-w-[200px] rounded-md h-32 bg-white p-4 shadow-md">
                    <div className="flex justify-between items-center ">
                    <h1 className="text-lg font-bold">{habit.name}</h1>
                    <p className='text-sm text-gray-500 '> {habit.goals[0].streak} day streak 🔥</p>
                    </div>
                   
            <div className="mt-3">
                <div className=" mt-2 w-full bg-gray-200 rounded-full h-2">
            <div
                className="bg-black h-2 rounded-full"
                style={{ width: `${percentage}%` }}
            />
        </div>    
             </div>
             <div className='mt-3 flex justify-between items-center'>
                <h3 className='text-sm'>
                    {percentage}% complete
                </h3>
                <button className=' px-4 py-2 border border-gray-300 rounded-md   text-sm  font-semibold text-black'>
                    Mark complete
                    </button>
                </div>
                   
                       
                </div>
                
            ))}
        </div>
    )
}

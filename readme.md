# Findr App

**_Findr_** is not just another web application. It's a unique platform designed to bridge the gap between students and placement opportunities, a crucial aspect of the Student Industrial Work Experience Scheme (SIWES).

With **_Findr_**, the process of finding and applying for industry placements becomes a breeze. Students can explore a wide range of opportunities and apply directly from their devices, while companies can easily manage and track applications.

## The Problem

Many students face significant challenges when seeking placements for the Student Industrial Work Experience Scheme (SIWES). Traditional methods often involve manual searches, lengthy application processes, and geographical constraints, making it difficult to secure suitable opportunities efficiently. Recruiters, on the other hand, struggle with managing numerous applications and communicating effectively with candidates.

### Benefits for Both Students and Recruiters

- _Time Saving_: Streamlined processes save time for both students and recruiters, allowing them to focus more on meaningful interactions and decision-making.
- _Improved Productivity_: Enhance productivity by reducing manual tasks associated with traditional placement processes.
- _Better Matching_: Facilitate better matching between students' skills and interests and companies' requirements, leading to more successful placements.

## Key Features

- Comprehensive Listings: Browse a diverse range of placement opportunities tailored to various fields of study, showcasing the breadth and depth of opportunities available from our valued partner companies.

- Effortless Application: Apply directly through the app using a streamlined application process.

- Company Insights: Access detailed profiles of partnering companies offering placements.

- Real-time Notifications: Stay updated with instant notifications on application statuses and new opportunities, powered by our robust and reliable concurrent mailing system.

- Personalized Dashboard: Manage applications and preferences efficiently through a personalized user dashboard.

## Why Choose Findr?

- Accessibility: Access placement opportunities from anywhere, eliminating geographical constraints.

- Efficiency: Streamline the search and application process to save valuable time.
Empowerment: Empower students by providing a variety of placement options across different industries.

## Technology & Architecture

- ### Client side

  - **Next.js**: For building the client-side, providing server-side rendering and a seamless user experience.

  - **Tailwind CSS**: For creating responsive, modern, and visually appealing UI components.

  - **React Hook Forms**: For managing form state and validation, ensuring a smooth user experience during the application process.

- ### Serverside

  - **Go** (Standard Library): Utilized for developing the server-side application, offering high performance and efficient concurrency handling.

  - **MongoDB Atlas**: A cloud-hosted NoSQL database service, used for storing and managing data securely and efficiently.

  - **Postman**: Employed for API documentation and testing, ensuring reliable and well-documented API endpoints.

## **System Architecture**

The architecture of Findr follows a modern, scalable design pattern.

- ### **Our Clientside**

 Built with Next.js, offering server-side rendering for faster page loads and better optimization.

- ### **Our Serverside**

 *Findr_is built with Go, a powerful language that ensures robust performance and scalability. The backend handles user authentication, data management, and business logic, providing a solid foundation for the app's functionality.

 The app adopts monolithic architecture
 for this project and suitable to provide a solution to the problem it solves, and the reason for this is that monolithic architecture is more manageable, simpler deployment, and lower initial complexity.

- ### **Database**

 MongoDB Atlas is used for data storage, providing flexibility and scalability for handling student and placement data.

- #### Database Design and Structure

  - **Collections**

        1. Users(Students & Companies): Stores user information, including name, email, password (hashed), and other profile details.

        2. Placements: Contains details about placement opportunities, such as company, description, and requirements.

        3. Applications: Tracks student applications, including status and related placement information.

## **API Documentation**

The API documentation for this application is managed using Postman, ensuring clear and comprehensive API documentation for development and testing.

This API will provide endpoints for user authentication, placement management, application tracking, and advanced search features using vector search and hopefully integrate a machine learning model algorithm for easy usage.

**Notice**: Detailed API documentation will be available on Postman, and it will be modified as we progress in this project.

### API Endpoints and Features

- **User Authentication**

  - **Sign Up**: Creating a student account and a company account that has available placement for students.

  - **Login**: Sign in student account and for companies that have registered.

- **Browse Opportunities**: Explore available placements.

- **Advanced Search and Filtering**: Search available placements based on your target company, location, or post.

- **Apply**: Submit applications directly through the web app.

- **Track Applications**: Monitor application status via the dashboard.

The full API documentation can be accessed through Postman.

[App documentation](Not available Yet)

## **About SIWES**

The Student Industrial Work Experience Scheme (SIWES) is an initiative aimed at bridging the gap between theoretical education and practical industry experience.

**_Findr_** supports this goal by connecting students with relevant placement opportunities and fostering practical skills development in a real-world environment.
